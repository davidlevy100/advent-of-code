using Common;

namespace Day00;

public class Day00Solution
{
    public void Run(string filename)
    {
        string path = Path.Combine("Day01", filename);
        var lines = Input.Lines(path);

        Console.WriteLine("Part 1: " + SolvePart1(lines));
        Console.WriteLine("Part 2: " + SolvePart2(lines));
    }

    private int SolvePart1(string[] lines)
    {
        // your logic here
        return 0;
    }

    private int SolvePart2(string[] lines)
    {
        // your logic here
        return 0;
    }
}
