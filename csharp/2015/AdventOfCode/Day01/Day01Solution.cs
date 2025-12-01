using Common;

namespace Day01;

public class Day01Solution
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
        int result = 0;

            foreach (string line in lines)
            {
                foreach (char c in line)
                {
                    if (c == '(')
                    {
                        result++;
                    }
                    else if (c == ')')
                    {
                        result--;
                    }
                }
            }

        return result;
    }

    private int SolvePart2(string[] lines)
    {
       {
        int result = 0;
        int counter = 0;

            foreach (string line in lines)
            {
                foreach (char c in line)
                {
                    if (c == '(')
                    {
                        result++;
                    }
                    else if (c == ')')
                    {
                        result--;
                    }

                    counter++;

                    if (result == -1)
                    {
                        return counter;
                    }

                }
            }

        return result;
    }
    }
}
