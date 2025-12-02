using Common;

namespace Day02;

public class Day02Solution
{
    public void Run(string filename)
    {
        string path = Path.Combine("Day02", filename);
        var lines = Input.Lines(path);

        Console.WriteLine("Part 1: " + SolvePart1(lines));
        Console.WriteLine("Part 2: " + SolvePart2(lines));
    }

    private ulong SolvePart1(string[] lines)
    {

        ulong result = 0;

        foreach (string line in lines)
        {
            string[] data = line.Split(",");

            foreach (string element in data)
            {
                string[] vals = element.Split("-");

                ulong start = ulong.Parse(vals[0]);
                ulong end = ulong.Parse(vals[1]);

                for (ulong i = start; i <= end; i++)
                {

                    string digits = i.ToString();

                    if (digits.Length % 2 != 0)
                    {
                        continue;
                    }

                    int half = digits.Length / 2;

                    if (digits[..half] == digits[half..])
                    {
                        result += i;
                    }

                }

            }

        }

        return result;

    }

    private ulong SolvePart2(string[] lines)
    {
        ulong result = 0;

        foreach (string line in lines)
        {
            string[] data = line.Split(',');

            foreach (string element in data)
            {
                string[] vals = element.Split('-');

                ulong start = ulong.Parse(vals[0]);
                ulong end = ulong.Parse(vals[1]);

                for (ulong i = start; i <= end; i++)
                {
                    string digits = i.ToString();

                    if (digits.Length < 2)
                        continue;

                    int len = digits.Length;

                    // Try every possible block size
                    for (int window = 1; window <= len / 2; window++)
                    {
                        // Block size must divide length
                        if (len % window != 0)
                            continue;

                        string model = digits[..window];
                        bool success = true;

                        // Compare each repeated block
                        for (int pos = window; pos < len; pos += window)
                        {
                            if (digits[pos..(pos + window)] != model)
                            {
                                success = false;
                                break;
                            }
                        }

                        if (success)
                        {
                            result += i;
                            break;  // count this number only once
                        }
                    }
                }
            }
        }

        return result;
    }

}
