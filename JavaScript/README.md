# more_commands_terminal.js
- [more\_commands\_terminal.js](#more_commands_terminaljs)
  - [Utilisation](#utilisation)
  - [Caractéristiques](#caractéristiques)
  - [Explication technique](#explication-technique)
  - [Known issue](#known-issue)
---
## Utilisation
* `node more_commands_terminal.js`
* taper: `help` puis regarder

## Caractéristiques
* Boucler de manière infinie (fait)
* Programme du système d'exploitation (voir Explication 2, "fait")
* Lister les processus en cours (en numérotant à partir de 1 dans le shell) avec la commande `lp` (voir Explication 2, "fait")
* Pouvoir tuer, mettre en pause ou reprendre un processus avec une commmande bing [-k|-p|-c] <processId> (problème)
* Exit shell on CTRL-P (ça marche avec CTRL-C)
## Explication technique

1. On utilise le package "commander" pour facilement créer les nouvelles commendes, les deux commendes "timeis" et "roll" démontrent les 2 façons de passer les paramètres.

2. Anisi on intègre les commands `Shell` ou `CMD` etc. selon l'OS. Par exemple, si on lance le programme JavaScript par `Command Prompt` de Microsoft, on peut utiliser les commandes de `Command Prompt`.
## Known issue
* Shell cmd exécute après Homemade cmd, n'importe comment Homemade cmd est exécuté (inconnu ou erreur).
* unknown option exception n'a pas traité, qui fait exit le programme.
* peux pas afficher "version" sans quitter