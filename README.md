# JSON Common Key Finder

This tool is designed to find common keys across multiple JSON files and save the results in an Excel file. It's useful for comparing JSON files to identify shared keys.

## Installation

Download the latest release for your platform from the [releases page](https://github.com/prabeshmagar/json-common-key-finder/releases). No additional installation steps are required; just download and run the executable.


## Usage

To use the Common Key Finder, you need to provide the JSON files as arguments. You can also specify the name of the output Excel file where the common keys will be saved. If not specified, the default output file name is `CommonKeys.xlsx`.

```bash
json-common-key-finder --output=name_of_file.xlsx file1.json file2.json ...`
```

## If no input files are provided, the tool will display a usage message:
``No input files provided. Usage: common-key-finder --output=name_of_file.xlsx file1.json file2.json ...``

## Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

