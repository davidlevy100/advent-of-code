using Common;

namespace Day04;

public class Day04Solution
{
    public void Run(string filename)
    {
        string path = Path.Combine("Day04", filename);
        var lines = Input.Lines(path);

        Console.WriteLine("Part 1: " + SolvePart1(lines));
        Console.WriteLine("Part 2: " + SolvePart2(lines));
    }

    private int SolvePart1(string[] lines)
    {
        char[,] grid = MakeMatrix(lines);
        return ScoreMatrix(grid).score;
    }

    private (int score, char[,] grid) ScoreMatrix(char[,] grid)
    {
        int rows = grid.GetLength(0);
        int cols = grid.GetLength(1);

        int score = 0;

        // track which cells to remove in this round
        bool[,] remove = new bool[rows, cols];

        (int dr, int dc)[] neighbors =
        [
            (-1, -1), (-1, 0), (-1, 1),
            ( 0, -1),          ( 0, 1),
            ( 1, -1), ( 1, 0), ( 1, 1),
        ];

        // ---------- PASS 1: decide who to remove ----------
        for (int r = 0; r < rows; r++)
        {
            for (int c = 0; c < cols; c++)
            {
                if (grid[r, c] != '@')
                    continue;

                int count = 0;

                foreach (var (dr, dc) in neighbors)
                {
                    int nr = r + dr;
                    int nc = c + dc;

                    if ((uint)nr < rows && (uint)nc < cols && grid[nr, nc] == '@')
                        count++;
                }

                if (count < 4)
                {
                    remove[r, c] = true;
                    score++;
                }
            }
        }

        // ---------- PASS 2: apply removals ----------
        char[,] newGrid = (char[,])grid.Clone();  // safe: new reference

        for (int r = 0; r < rows; r++)
        {
            for (int c = 0; c < cols; c++)
            {
                if (remove[r, c])
                    newGrid[r, c] = '.';
            }
        }

        return (score, newGrid);
    }

    private int SolvePart2(string[] lines)
    {
        char[,] grid = MakeMatrix(lines);
        int total = 0;

        while (true)
        {
            var (score, next) = ScoreMatrix(grid);
            if (score == 0)
                break;

            total += score;
            grid = next;
        }

        return total;
    }

    private static char[,] MakeMatrix(string[] lines)
    {
        char[,] result = new char[lines.Length, lines[0].Length];

        for (int r = 0; r < lines.Length; r++)
            for (int c = 0; c < lines[r].Length; c++)
                result[r, c] = lines[r][c];

        return result;
    }
}
