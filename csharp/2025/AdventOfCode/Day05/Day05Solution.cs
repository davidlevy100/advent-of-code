using Common;

namespace Day05;

public class Day05Solution
{
    public void Run(string filename)
    {
        string path = Path.Combine("Day05", filename);
        var lines = Input.Lines(path);

        Console.WriteLine("Part 1: " + SolvePart1(lines));
        Console.WriteLine("Part 2: " + SolvePart2(lines));
    }

    private int SolvePart1(string[] lines)
    {
        int result = 0;

        var ranges = new List<(ulong start, ulong end)>();

        foreach (string line in lines)
        {
            if (string.IsNullOrWhiteSpace(line)) continue;

            if (line.Contains('-')) {
                string[] vals = line.Split('-');

                ulong start = ulong.Parse(vals[0]);
                ulong end = ulong.Parse(vals[1]);

                ranges.Add((start, end));
            } else
            {
                ulong thisNum = ulong.Parse(line);

                foreach (var (start, end) in ranges)
                {
                    if (thisNum >= start && thisNum <= end)
                    {
                       result++;
                       break; 
                    } 
                }
            }
        }

        return result;
    }

    private ulong SolvePart2(string[] lines)
    {

        var ranges = new List<(ulong start, ulong end)>();

        // --- Parse ranges ---
        foreach (string line in lines)
        {
            if (string.IsNullOrWhiteSpace(line)) 
                break;

            string[] parts = line.Split('-');
            ulong start = ulong.Parse(parts[0]);
            ulong end   = ulong.Parse(parts[1]);

            ranges.Add((start, end));
        }

        // --- Sort by start ---
        ranges.Sort((a, b) => a.start.CompareTo(b.start));

        // --- Merge ---
        var merged = new List<(ulong start, ulong end)>();

        foreach (var r in ranges)
        {

            if (merged.Count == 0)
            {
                merged.Add(r);
                continue;
            }

            var last = merged[^1];

            if (r.start > last.end + 1)
            {
                // no overlap
                merged.Add(r);
            }
            else
            {
                // overlap → extend the last interval
                merged[^1] = 
                    (last.start, Math.Max(last.end, r.end));
            }
        }

        ulong result = 0;

        foreach (var m in merged)
        {
            result += (m.end - m.start + 1);
        }

        return result;

    }
}
