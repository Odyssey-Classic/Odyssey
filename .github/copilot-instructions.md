# Copilot Coding Instructions

- Never use log.Fatal, log.Fatalf, log.Fatalln, or os.Exit outside the main package. Instead, return errors and let main handle termination.
- Please use the stretchr testify suite when setting up tests.
