using Common;

namespace Day01;

public class Day01Solution
{
    public void Run(string filename)
    {
        // Build path (Day01/input.txt or test.txt)
        string path = Path.Combine("Day01", filename);

        // Load all lines
        var lines = Input.Lines(path);

        Console.WriteLine("Part 1: " + SolvePart1(lines));
        Console.WriteLine("Part 2: " + SolvePart2(lines));
    }

    private int SolvePart1(string[] lines)
    {
        int dial = 50;     // starting position
        int zeroes = 0;    // times we end a rotation on 0

        foreach (string line in lines)
        {
            char direction = line[0];           // 'L' or 'R'
            int magnitude = int.Parse(line[1..]); // parse the number after the direction

            // Apply the rotation
            dial += (direction == 'L' ? -magnitude : magnitude);

            // Wrap into [0–99]
            dial %= 100;
            if (dial < 0) dial += 100;

            // Count if the *final position* lands on 0
            if (dial == 0)
                zeroes++;
        }

        return zeroes;
    }

    private int SolvePart2(string[] lines)
    {
        int dial = 50;     // starting position
        int zeroes = 0;    // count *each click* that lands on 0

        foreach (string line in lines)
        {
            char direction = line[0];             // 'L' or 'R'
            int magnitude = int.Parse(line[1..]); // how many clicks to rotate

            int step = direction == 'L' ? -1 : 1; // step per click

            // Simulate click-by-click movement
            for (int i = 0; i < magnitude; i++)
            {
                dial += step;          // move one click

                // Wrap into [0–99]
                dial %= 100;
                if (dial < 0) dial += 100;

                // Count every click that *lands* on 0
                if (dial == 0)
                    zeroes++;
            }
        }

        return zeroes;
    }
}
