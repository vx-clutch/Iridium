# Grammar

## Hierarchy for Iridium
$$
\text{Program} \to \text{List of statements} \to \text{block} \to \left\{ 
\begin{matrix}
\text{operations} \\
\text{literals} \\
\text{variables}
\end{matrix} 
\right\}
$$

## Hierarchy for Onyx

$$
\text{Program} \to \text{Main function} \to \left\{
    \begin{matrix}
    \text{Functions} \\
    \left\{ 
\begin{matrix}
\text{operations} \\
\text{literals} \\
\text{variables}
\end{matrix} 
\right\} \to \text{Function}
\end{matrix}
\right\}
$$

# Syntax Examples
These examples exit with code ```20```
## Examples for Iridium
```
fn main = {
    exit = 20
} => void
```
## Examples for Onyx
```
MAIN:
    EXIT, 20 
```