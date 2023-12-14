import { expect } from 'chai';
import { toSignificant, validate } from '../src/precision';

describe('Precision Tests', function () {
    describe('toSignificant', function () {
        const tests = [
            { input: "0.0", expect: "0", },
            { input: "0.0000000000000004", expect: "0.0000000000000004", },
            { input: "0.0001", expect: "0.0001", },
            { input: "000.0001", expect: "0.0001", },
            { input: "0.100103456123", expect: "0.10010345", },
            { input: "0.012345678999", expect: "0.012345678", },
            { input: "0.001234023499", expect: "0.0012340234", },
            { input: "0.012345678928", expect: "0.012345678", },
            { input: "0.100001023499", expect: "0.10000102", },
            { input: "0.012345", expect: "0.012345", },
            { input: "1.0000000234", expect: "1", },
            { input: "1.0002", expect: "1.0002", },
            { input: "2", expect: "2", },
            { input: "1000000060", expect: "1000000000", },
            { input: "1000000060.123", expect: "1000000000", },
            { input: "12345", expect: "12345", },
            { input: "222.222222", expect: "222.22222", },
            { input: "2222022.2202", expect: "2222022.2", },
            { input: "218166.0002", expect: "218166", },
        ];

        tests.forEach(test => {
            it(`should truncate ${test.input} to ${test.expect}`, function () {
                expect(toSignificant(test.input, 8)).to.equal(test.expect);
            });
        });
    });

    describe('validate', function () {
        it('should return error on negative numbers', function () {
            expect(() => validate("-12345", 18)).to.throw();
        });

        it('should return error if precision is greater than allowed', function () {
            expect(() => validate("0.00000000000000012345678", 18)).to.throw();
        });

        it('should pass for valid input', function () {
            expect(() => validate("1.2345", 18)).not.to.throw();
        });
    });
});

