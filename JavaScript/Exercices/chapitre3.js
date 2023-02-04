function countBs(chaine) {
  let compteur = 0;
  for(let i = 0; i<chaine.length; i++){
    if(chaine[i] == 'B'){
      compteur++;
    }
  }
  return compteur;
}

function countChar(chaine,character) {
  let compteur = 0;
  for(let i = 0; i<chaine.length; i++){
    if(chaine[i] == character){
      compteur++;
    }
  }
  return compteur;
}
