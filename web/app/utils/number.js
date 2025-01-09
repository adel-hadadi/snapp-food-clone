export function convertEnToFaNumber(input) {
  input = input / 10;
  const numberStr = input.toString();

  // Replace English digits with Persian digits
  const persianDigits = ["۰", "۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹"];
  const englishToPersian = (number) =>
    number.replace(/\d/g, (digit) => persianDigits[parseInt(digit, 10)]);

  // Add comma separators for every 3 digits
  const addCommas = (number) => number.replace(/\B(?=(\d{3})+(?!\d))/g, "،"); // Use Persian comma (\u060C)

  // Convert to Persian digits and format with commas
  const parts = numberStr.split("."); // Handle decimals if present
  parts[0] = addCommas(parts[0]); // Add commas to the integer part
  const formattedNumber = parts.join(".");

  return englishToPersian(formattedNumber);
}
