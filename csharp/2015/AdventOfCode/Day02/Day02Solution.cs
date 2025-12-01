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

    private int SolvePart1(string[] lines)
    {
        int result = 0;

        foreach (string line in lines)
        {
            string[] parts = line.Split('x');
            int l = int.Parse(parts[0]);
            int w = int.Parse(parts[1]);
            int h = int.Parse(parts[2]);

            int side1 = l * w;
            int side2 = w * h;
            int side3 = h * l;

            int smallestSide = Math.Min(side1, Math.Min(side2, side3));

            result += 2 * side1 + 2 * side2 + 2 * side3 + smallestSide;

        }

        return result;

    }

    private int SolvePart2(string[] lines)
    {
        int result = 0;

        foreach (string line in lines)
        {
            string[] parts = line.Split('x');

            int l = int.Parse(parts[0]);
            int w = int.Parse(parts[1]);
            int h = int.Parse(parts[2]);

            int[] dimensions = [l, w, h];

            Array.Sort(dimensions);

            int bow = l * w * h;
            int ribbon = bow + 2 * (dimensions[0] + dimensions[1]);

            result += ribbon;
        }

        return result;
    }
}
