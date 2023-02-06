import os
import random

def get_matrix_sizes():
    while True:
        rows_A = int(input("Enter number of rows in matrix A: "))
        cols_A = int(input("Enter number of columns in matrix A: "))
        rows_B = int(input("Enter number of rows in matrix B: "))
        cols_B = int(input("Enter number of columns in matrix B: "))
        if cols_A == rows_B:
            break
        print("Matrix multiplication not possible. Try again.")
    return rows_A, cols_A, rows_B, cols_B

def get_range():
    a = int(input("Enter starting number of range: "))
    b = int(input("Enter ending number of range: "))
    return a, b

def generate_matrices(rows_A, cols_A, rows_B, cols_B, a, b):
    A = [[random.randint(a, b) for j in range(cols_A)] for i in range(rows_A)]
    B = [[random.randint(a, b) for j in range(cols_B)] for i in range(rows_B)]
    return A, B

def save_matrices(A, B, pid):
    with open(f"{pid}_A.txt", "w") as f:
        for row in A:
            f.write(",".join(str(x) for x in row) + "\n")
    with open(f"{pid}_B.txt", "w") as f:
        for row in B:
            f.write(",".join(str(x) for x in row) + "\n")

def matrix_multiplication(A, B):
    result = [[0 for j in range(len(B[0]))] for i in range(len(A))]
    for i in range(len(A)):
        for j in range(len(B[0])):
            for k in range(len(B)):
                result[i][j] += A[i][k] * B[k][j]
    return result

def save_result(result, pid):
    with open(f"{pid}_AB.txt", "w") as f:
        for row in result:
            f.write(",".join(str(x) for x in row) + "\n")

if __name__ == "__main__":
    pid = str(os.getpid())
    rows_A, cols_A, rows_B, cols_B = get_matrix_sizes()
    a, b = get_range()
    A, B = generate_matrices(rows_A, cols_A, rows_B, cols_B, a, b)
    save_matrices(A, B, pid)
    result = matrix_multiplication(A, B)
    save_result(result, pid)