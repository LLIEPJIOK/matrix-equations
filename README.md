# Matrix equations 

## Project Description  

This program solves a system of linear equations using two methods:  
1. **Gauss method**:  
   It solves the system using the Gauss method with the choice of the leading element by row ans finds the inverse matrix. 

2. **Reflection (Householder) method**:  
   It decomposes the matrix into an orthogonal matrix (Q) and an upper triangular matrix (R) to solve the system via matrix factorization.

The system is represented by the equation:  
\[
A \cdot x = b
\]  
Where:  
- The matrix **A** is defined as:  
  \( a_{ii} = 5i, \quad a_{ij} = -(i + \sqrt{j}) \) for \( i \neq j \),  
  with \( i, j = 1, N \), where \( N = 15 \).
  
- The vector **b** is defined as:  
  \( b_i = 3\sqrt{i} \) for \( i = 1, N \).

The matrix **A** and vector **b** are created dynamically based on these formulas.

---  

## Installation  

1. Clone the repository:  
   ```bash  
   git clone git@github.com:LLIEPJIOK/matrix-equations.git 
   ```  

2. Navigate to the project folder:  
   ```bash  
   cd matrix-equations 
   ```  

3. Run the program:  
	3.1. Run Gauss method
   ```bash  
   go run cmd/gauss/main.go  
   ```  

	3.2. Run Householder method
   ```bash  
   go run cmd/householder/main.go  
   ```  