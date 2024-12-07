
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
  - **Phone Masking**: Masks phone numbers, keeping the first three digits and masking the rest.

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

### Supported Anonymization Techniques

| Technique      | Description                                                                 |
|----------------|-----------------------------------------------------------------------------|
| `mask`         | Replace data with partial information or asterisks.                        |
| `pseudonymize` | Replace data with random placeholders (e.g., "Person 1").                  |
| `hash`         | Apply irreversible hashing (e.g., SHA-256).                                |
| `generalize`   | Replace numeric data with broader categories (e.g., 29 → "20-29").         |
| `phone_mask`   | Keep the first three digits of a phone number, masking the rest (e.g., `555-****`). |

### Example Command
```bash
./data-anonymizer --input test_data.csv --columns "Name,Email,Phone" --techniques "mask,hash,phone_mask"
```

---

#### Example Input
```csv
Name,Email,Age,Phone
John Doe,john.doe@example.com,29,555-1234
Jane Smith,jane.smith@example.com,34,555-5678
```

---

#### Example Output
```csv
Name,Email,Age,Phone
Jo******,836f82db99121b3481011f16b49dfa5fbc714a0d1b1b9f784a1ebbbf5b39577f,29,555-****
Ja********,f2d1f1c853fd1f4be1eb5060eaae93066c877d069473795e31db5e70c4880859,34,555-****
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

