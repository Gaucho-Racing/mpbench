def print_binary(num):
    """Print a number in binary format with spacing between bytes."""
    binary = format(num, '016b')  # 16-bit binary representation
    return f"{binary[:8]} {binary[8:]}"  # Space between bytes

def main():
    while True:
        try:
            # Get user input
            num = int(input("\nEnter a number (0-65535) or -1 to quit: "))
            
            if num == -1:
                break
                
            if not 0 <= num <= 65535:
                print("Please enter a number between 0 and 65535")
                continue
            
            # Convert to bytes in both endian formats
            big_endian = num.to_bytes(2, byteorder='big')
            little_endian = num.to_bytes(2, byteorder='little')
            
            # Print detailed representation
            print("\n=== Number Analysis ===")
            print(f"Decimal: {num}")
            print(f"Binary:  {print_binary(num)}")
            print("\nBig Endian (network order):")
            print(f"Bytes:   {big_endian.hex(' ')}")
            print(f"Binary:  {' '.join(format(b, '08b') for b in big_endian)}")
            
            print("\nLittle Endian:")
            print(f"Bytes:   {little_endian.hex(' ')}")
            print(f"Binary:  {' '.join(format(b, '08b') for b in little_endian)}")
            
            # Show reconstruction
            print("\nReconstruction:")
            print(f"From Big Endian:    {int.from_bytes(big_endian, byteorder='big')}")
            print(f"From Little Endian: {int.from_bytes(little_endian, byteorder='little')}")
            
        except ValueError:
            print("Please enter a valid number")
        except Exception as e:
            print(f"An error occurred: {e}")

if __name__ == "__main__":
    print("Welcome to the Endianness Explorer!")
    print("This tool helps visualize how 2-byte integers are stored in different byte orders")
    main()
