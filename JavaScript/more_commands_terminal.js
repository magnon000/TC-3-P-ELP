const readline = require('readline');
const program = require('commander');  // process.argv[2] if we don't use commander Lib
const {exec} = require('child_process');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

let versionTriggered = false;
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
    .description('lancer les dés de Dn (exemple. D20: 1-20, CMD: roll 20')
    .action((arg1) => {
        const randomNumber = Math.floor(Math.random() * arg1) + 1;
        if (randomNumber === 1) {
            console.log("Total failure!");
        } else if (randomNumber == arg1){  // ignore warning
            console.log("Huge success!");
        } else {
            console.log(randomNumber);
        }
    });

// program.on('--help', () => {
//     console.log('\nCommands possibles:');
//     program.commands.forEach((cmd) => {
//         console.log(`    ${cmd._name} - ${cmd._description}`);
//     });
// });

program.on('command:*', () => {
    console.error('HOMEMADE command inconnue: %s\n      taper "help" pour voir la liste de toutes les commands.', program.args.join(' '));
    prompt();
});

program.unknownOption = (flag) => {
    console.error(`Unknown option: ${flag}`);
    prompt();
};

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

        if (versionTriggered) {
            versionTriggered = false;
            prompt();
            return;
        }

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

        // if (!program.commands.some((cmd) => cmd._name === command)) {
        //     console.log(`command inconnue: ${command}`);
        //     console.log('Utiliser --help for voir tout');
        //     prompt();
        // }

        // const helpIndex = args.indexOf('--help');
        // if (helpIndex !== -1) {
        //     args.splice(helpIndex, 1);
        //     program.parse(['node', 'index.js', command, ...args, '--help']);
        //     prompt();
        //     return;
        // }

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
