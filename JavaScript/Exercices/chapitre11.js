
//ne fonctionne pas : le site ne reconnait pas sa propre fonction readStorage
async function locateScalpel(nest) {
  let nest_scalpel = await anyStorage(nest, nest.name, "scalpel");
  let scalpel_found = nest_scalpel == nest.name
  while(!scalpel_found) {
    let previous = nest_scalpel
    nest_scalpel = await anyStorage(nest_scalpel, nest_scalpel.name, "scalpel");
    scalpel_found = nest_scalpel == previous
    if (scalpel_found) {
     return nest_scalpel
    }
  }
}

function locateScalpel2(nest) {
  new_nest = anyStorage(nest, nest.name, "scalpel")
  if (nest.name == new_nest) {
    return nest
  } else {
    locateScalpel2(new_nest)
  }
}


