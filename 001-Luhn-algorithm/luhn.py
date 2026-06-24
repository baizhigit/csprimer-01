
def verify_imperative(digits):
    total = 0
    for i, d in enumerate(reversed(digits)):
        x = int(d) * (1 + i % 2)
        total += x // 10 + x
    return total % 10 == 0


def luhn_validate(card_number):
    total = 0
    # Remove spaces/hyphens if present
    digits = [int(d) for d in str(card_number) if d.isdigit()]
    
    if not digits:
        raise ValueError("No digits found in input")
    
    # Double every second digit from right
    for i in range(len(digits) - 2, -1, -2):
        doubled = digits[i] * 2
        digits[i] = doubled - 9 if doubled > 9 else doubled
    
    return sum(digits) % 10 == 0


def luhn_verify_check(card_number):
    digits = [int(d) for d in str(card_number) if d.isdigit()]
    
    if len(digits) < 2:
        return False
    
    arr = [0, 2, 4, 6, 8, 1, 3, 5, 7, 9]
    
    # Separate check digit
    check, *rest = reversed(digits)  # Now check is last digit, rest are others reversed
    
    # Process the rest (which are already reversed)
    total = 0
    for i, t in enumerate(rest):
        total += arr[int(t)] if i % 2 == 0 else int(t)
    
    return 10 - (total % 10) == int(check)


def test_luhn():
    test_cases = [
        ("17893729974", True),
        ("4539 1488 0343 6467", True),  # Valid Visa
        ("6011 1111 1111 1117", True),  # Valid Discover
        ("1234 5678 9012 3456", False),  # Invalid
        ("", False),  # Empty
        ("5", False),  # Too short
    ]
    
    for input_val, expected in test_cases:
        try:
            result = luhn_validate(input_val)
            assert result == expected, f"Failed for {input_val}"
        except ValueError:
            assert expected is False, f"Should not raise for {input_val}"
    print("All tests passed!")


if __name__ == '__main__':
    assert verify_imperative("17893729974")
    assert not verify_imperative("17893729975")
    # test_luhn()
    print("ok!")