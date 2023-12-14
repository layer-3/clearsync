export function toSignificant(input: string, sigDigits: number): string {
    let output = '';
    let outputSig = 0;
    let dotFound = false;

    for (const char of input) {
        if (outputSig === sigDigits) break;

        if (char === '0' && outputSig === 0) {
            if (output.length === 0 || dotFound) {
                output += char;
            }
            continue;
        } else if (char === '.') {
            output += char;
            dotFound = true;
            continue;
        }

        if (char !== '0' || outputSig > 0) {
            outputSig++;
        }
        output += char;
    }

    return appendZeros(input, output);
}

function appendZeros(input: string, output: string): string {
    const dotIndex = input.indexOf('.');
    if (dotIndex != -1 && output.length < dotIndex) {
        const zerosToAppend = dotIndex - output.length;
        output += '0'.repeat(zerosToAppend);
    } else if (dotIndex == -1) {
        const zerosToAppend = input.length - output.length;
        output += '0'.repeat(zerosToAppend);
    } else {
        // Remove trailing zeros and decimal point if the output ends with them
        output = output.replace(/\.0+$/, '');
    }

    return output;
}

export function validate(input: string, maxPrecision: number) {
    if (input.startsWith('-')) {
        throw new Error("Input must not be negative");
    }

    const dotIndex = input.indexOf('.');
    const decimalPlaces = dotIndex !== -1 ? input.length - dotIndex - 1 : 0;
    if (decimalPlaces > maxPrecision) {
        throw new Error(`Input must not exceed max precision (allowed: ${maxPrecision}, actual: ${decimalPlaces})`);
    }
}
