const readline = require('readline');
const program = require('commander');  // process.argv[2] if we don't use commander Lib
const {exec} = require('child_process');
const os = require('os');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

// let versionTriggered = false;
// let hmCmdFlag = false;

program
    // .version('1.0', '-v, --version', 'version 1.0')
    .version("1.0")
    .description('------timeis: afficher la date, le temps comme on veut')
    .command('timeis')
    // .option('-v --version', 'afficher la version')
    .option('-d --date', 'afficher la date')
    .option('-t --time', 'afficher le temps')
    .action((commander) => {
        // if (commander.version) {
        //     console.log(program.version());
        //     return;
        // }
        const date = new Date()
        if (commander.date) {
            console.log(`${date.getDate()}-${date.getMonth() + 1}-${date.getFullYear()}`)
            return;
        }
        if (commander.time) {
            console.log(`${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`)
            return;
        }
        console.log(`${date.getDate()}-${date.getMonth() + 1}-${date.getFullYear()}  ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`)
    });

program
    .command('roll <arg1>')
    .version("1.0")
    .description('------lancer les dés de Dn (exemple. D20: 1-20, CMD: roll 20')
    .action((arg1) => {
        console.log("Lancer un dé:")
        const randomNumber = Math.floor(Math.random() * arg1) + 1;
        if (randomNumber === 1) {
            console.log("Total failure!");
        } else if (randomNumber == arg1) {  // ignore warning
            console.log("Huge success!");
        } else {
            console.log(randomNumber);
        }
    });

program
    .command('lp')
    .version("1.0")
    .description('------lister tous processus')
    .action(() => {
        if (os.platform() === 'linux') {
            exec('ps -A', (err, stdout, stderr) => {
                if (err) {
                    console.error(`\nErreur: ${err}`);
                    prompt();
                    return;
                }
                let processes = stdout.split('\n');
                processes.shift();  // remove header row
                processes = processes.map((p, index) => `${index + 1}. ${p}`);
                console.log(processes.join('\n'));
                prompt();
            });
        }
        if (os.platform() === 'win32') {
            exec('tasklist', (err, stdout, stderr) => {
                if (err) {
                    console.error(`\nErreur: ${err}`);
                    prompt();
                    return;
                }
                let processes = stdout.split('\n');
                processes.shift();  // remove header row
                processes = processes.map((p, index) => `${index - 1}. ${p}`);
                console.log(processes.join('\n'));
                prompt();
            });
        }
    });

program
    .command("bing [options]")
    .version("1.0")
    .description("------tuer, mettre en pause ou reprendre un processus")
    .option("-k <processId>, --kill <processId>", "tuer un processus")
    .option("-p <processId>, --pause <processId>", "mettre en pause un processus, que sur Linux")
    .option("-c <processId>, --continue <processId>", "reprendre un processus, que sur Linux")
    .action((options, processId) => {
        const processMap = new Map();
        if (!processId) {
            console.error("Erreur: 'processId' manquant");
            return;
        }
        const childProcess = processMap.get(processId);
        // const childProcess = processMap.get(parseInt(processId)); // not working: Processus non trouvé: [object Object]
        if (!childProcess) {
            console.log(processId)
            console.error(`Erreur: Processus non trouvé: ${processId}`);
            return;
        }
        if (options.kill) {
            console.log(`Tuer: ${(processId)}`);
            childProcess.kill();
            processMap.delete((processId));
            console.log(`Tué: ${processId}`);
        } else if (options.pause) {
            console.log(`Mettre en pause: ${processId}`);
            childProcess.stdin.pause();
            console.log(`Mis en pause: ${processId}`);
        } else if (options.continue) {
            console.log(`Reprendre: ${processId}`);
            childProcess.stdin.resume();
            console.log(`Repris: ${processId}`);
        } else {
            console.error("Erreur: action non valide");
        }
    });


program.on('command:*', () => {
    console.error('HOMEMADE commande inconnue: %s\n      taper "help" pour voir la liste de toutes les commandes.', program.args.join(' '));
    prompt();
});

program.unknownOption = (flag) => {
    console.error(`Unknown option: ${flag}`);
    prompt();
};

rl.on("SIGINT", function () {
    console.log("Exit on CTRL+C...");
    process.exit();
});

// program.on('version', () => {
//     versionTriggered = true;
//     console.log(`v${program.version()}`);
// });

const prompt = () => {
    rl.question('$ ', (input) => {
        // handle entering "Enter"
        if (!input) {
            prompt();
            return;  // use return to avoid trigger invalid command check
        }

        // if (versionTriggered) {
        //     versionTriggered = false;
        //     prompt();
        //     return;
        // }

        const [command, ...args] = input.split(' ');
        // handle manual exit
        if (command === 'exit') {
            console.log('Terminal terminé !');
            rl.close();
            return;
        }
        // handle help to show all commands
        if (command === 'help') {
            console.log(program.helpInformation());
            prompt();
            return;
        }

        // program.parse([process.argv[0], process.argv[1], command, ...args]);
        // if (program.parse(['node', 'more_commands_terminal.js', command, ...args])) {
        //     prompt();
        //     return;
        // }
        program.parse(['node', 'more_commands_terminal.js', command, ...args])
        exec(input, (err, stdout, stderr) => {
            if (err) {
                console.error(`\nShell/CMD command inconnue: ${err}`);
                prompt();
                return;
            }
            console.log(`${stdout}`);
            console.log(`Shell/CMD message d'erreur (s'il existe):\n ${stderr}`);
            prompt();
        });
    });
};

console.log('More-commands terminal: (taper "exit" pour terminer, "help" pour toutes commands');
console.log('                                               ou utiliser les commands Shell/CMD)');
prompt();
