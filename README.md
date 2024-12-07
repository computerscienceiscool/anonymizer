
# Data Anonymizer

The **Data Anonymizer** is a command-line tool written in Go that anonymizes sensitive data in CSV files. It supports various anonymization techniques for specified columns and generates a new CSV file with anonymized data.

---

## Features

- **Input CSV Support**: Accepts a CSV file with column headers as input.
- **Anonymization Techniques**:
  - **Masking**: Replaces sensitive data with partial data or asterisks.
  - **Pseudonymization**: Replaces data with random placeholders.
  - **Hashing**: Applies irreversible hashing (e.g., SHA-256).
  - **Generalization**: Broadens specific values into categories.
- **Customizable**: Specify columns and techniques via command-line arguments.
- **Logging**: Logs anonymization activities for transparency.
- **Security**: Avoids storing sensitive data in memory longer than necessary.

---

## Installation

### Prerequisites
- **Go** installed (version 1.20+).
- A Linux environment with a terminal (works on macOS and Windows as well).

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/data-anonymizer.git
   cd data-anonymizer
   ```
2. Build the executable:
   ```bash
   go build -o data-anonymizer main.go
   ```

---

## Usage

### Command-Line Arguments
| Argument       | Description                                               |
|----------------|-----------------------------------------------------------|
| `--input`      | Path to the input CSV file.                               |
| `--columns`    | Comma-separated list of columns to anonymize.             |
| `--techniques` | Comma-separated list of anonymization techniques.         |
| `--help`       | Displays usage information.                               |

### Example
```bash
./data-anonymizer --input test_data.csv --columns "Name,Email,Phone" --techniques "mask,hash,pseudonymize"
```

#### Input (test_data.csv)
```csv
Name,Email,Age,Phone
John Doe,john.doe@example.com,29,555-1234
Jane Smith,jane.smith@example.com,34,555-5678
```

#### Output (test_data_output.csv)
```csv
Name,Email,Age,Phone
Jo********,e3b0c44298fc1c149afbf4c8996fb924,29,Phone 1
Ja********,c81e728d9d4c2f636f067f89cc14862c,34,Phone 2
```

---

## Development

### Adding Anonymization Techniques
New anonymization techniques can be added by modifying the `anonymize` function in `main.go`. Ensure each technique is properly tested.

---

## Future Improvements
- Support for additional file formats (e.g., Excel, JSON).
- Interactive mode for column selection and technique assignment.
- Batch processing for multiple files.

---

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.
