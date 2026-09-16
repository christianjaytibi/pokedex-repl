# Pokedex REPL

> A **Pokedex** built in Go. It queries the [PokeAPI](https://pokeapi.co/), implements a custom thread-safe caching layer, and runs an interactive REPL.

---

## Table of Contents
- [About The Project](#about-the-project)
- [Features & Commands](#features--commands)
- [Catch Rate Formula](#catch-rate-formula)
- [Prerequisites](#prerequisites)
- [Installation](#installation--getting-started)
- [Usage Walkthrough](#usage-walkthrough)

---

## About The Project

This project emphasizes writing idiomatic Go without external web frameworks or heavy libraries. It focuses on low-level backend fundamentals: parsing command line inputs, issuing HTTP client requests, serializing and deserializing JSON payloads, managing state, and implementing concurrent data protection via mutexes.

---

## Features & Commands

* **REPL Engine:** Dispatch matching callback functions mapped in a registry.
* **Thread-Safe Caching:** A custom caching implementation using a Go map and a `sync.Mutex`. A background goroutine purges expired keys safely.
* **State Management:** Uses a shared configuration struct passed by pointer across commands.

The CLI exposes a custom interactive REPL prompt supporting the following commands:

| Command | Arguments | Description |
| :--- | :--- | :--- |
| **`help`** | None | Displays the help menu and lists all available commands. |
| **`exit`** | None | Exits the Pokedex REPL session. |
| **`map`** | None | Fetches and displays the next 20 location areas in the Pokémon world. Results are cached. |
| **`mapb`** | None | Fetches and displays the previous 20 location areas. |
| **`explore`** | `<location_name>` | Lists all Pokémon encountered within a specific location area. |
| **`catch`** | `<pokemon_name>` | Attempts to catch a Pokémon based on its base experience level, adding it to your Pokedex on success. |
| **`inspect`** | `<pokemon_name>` | Displays details (height, weight, stats, types) for any Pokémon you have successfully caught. |
| **`pokedex`** | None | Lists the names of all Pokémon currently stored in your personal collection. |

---
## Catch Rate Formula

$$
P = \text{Catch Probability (\\%)} =
\max\left(5, \min\left(95, \left(105 - \frac{\text{Base EXP}}{3}\right)\right)\right)
$$

### Decision Rule

Generate a random number $R$ between $0.0$ and $1.0$.

- If $R \le P$, where $P$ is the **Catch Probability (\%)**, the catch is successful.
- If $R > P$, the Pokémon breaks free or flees.

---

## Prerequisites

* **Go** (version 1.27 or higher)

---

## Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/christianjaytibi/pokedex-repl.git
    cd pokedex-repl
   ```

---

## Usage Walkthrough
* **Run directly:**
  ```bash
  go run .
  ```

* **Build a binary and run:**
  ```bash
  go build .
  ./pokedex-repl
  ```
