# Thrive Property Calculator

## 🚀 How to use

1. Run `make build` to compile the application. It will result in binary file `bin/thrivepropcalc`.
2. Prepare the inputs in `input.txt`
3. Run `make run` to run the binary file. You should provide an `input.txt`, otherwise the app won't run as intended.
4. Run `make clean` to clean up the compiled file.
5. Alternatively, you can build and run the application with a single command `make all`

## ❓ How does it work?

First, we need to understand the rules.

1. Base property value is calculated per square meter:
   - Standard location: Rp 10,000,000/m^2
   - Premium location: Rp 15,000,000/m^2

2. Annual value adjustments:
   - Land appreciates 5% annually
   - Residential buildings depreciate 2.5% annually
   - Commercial buildings depreciate 3.5% annually

3. Location bonuses:
   - Premium location adds 20% to total value
   - Corner lot adds 15% to total value

5. Monthly maintenance calculation:
   - Residential base fee: Rp 2,500/m²
   - Commercial base fee: Rp 3,500/m²
   - Security fee: Rp 1,000/m²
   - Cleaning fee: Rp 800/m²

The system maintains two monitoring systems. Only the most recent values are recorded:
   - Property Value Monitor
   - Maintenance Fee Calculator