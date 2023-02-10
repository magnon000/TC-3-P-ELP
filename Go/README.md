# Golang parallel Matrix multiplication (server-client TCP)
- [Golang parallel Matrix multiplication (server-client TCP)](#golang-parallel-matrix-multiplication-server-client-tcp)
	- [Usage](#usage)
		- [1. mode input](#1-mode-input)
		- [2. mode file](#2-mode-file)
		- [3. exception handle](#3-exception-handle)
---
## Usage
### 1. mode input
> * Manual input all numbers, number is seperated by `,` , line is seperated by `Enter`. 
> * Use `/` to end input of one matrix.
> * Use `end/` to end input.
> 
Example (multiply 3 matrix):
```
\#       N.1 matrix:
1,2,3
1,2,3/
\#       N.2 matrix:
1,2
-1,-2
1,2/
\#       N.3 matrix:
1,0
0,1/
\#       N.4 matrix:
end/
```
### 2. mode file
> * Need to define file name (path+partial name) for multiplication result. File name automatically contains timestamp.
```
resultFile     = "./res_"
```
> * Need to specify input file names, and the number of files.
> * For now we used a simple list []string to manager matrix files (number of files < 7).

Example:
```
matriceA_raw   = "matrix_input_txt/test_set1/A_1000x500.txt"
...
matriceNbr     = 7 
var matrix_raw_list = [...]string{ // TODO: read .txt names from a file?
	matriceA_raw,
	matriceB_raw,
	matriceC_raw,
	matriceD_raw,
	matriceE_raw,
	matriceF_raw,
	matriceG_raw,
	matriceFin_raw}
```
> * content in input file use following format: (file ends without empty lines)
> 
```
1,-2.1,-3
-1,-2,3.3
``` 
### 3. exception handle
If one step of multiplication is not possible, previous multiplication result will be saved. (or 1st matrix if 1st multiplication is already not doable)
