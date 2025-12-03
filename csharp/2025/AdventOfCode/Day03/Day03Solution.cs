using Common;

namespace Day03;

public class Day03Solution
{
    public void Run(string filename)
    {
        string path = Path.Combine("Day03", filename);
        var lines = Input.Lines(path);

        Console.WriteLine("Part 1: " + SolvePart1(lines));
        Console.WriteLine("Part 2: " + SolvePart2(lines));
    }

    // ------------------------------------------------------------
    // PART 1
    // Find the largest possible TWO-DIGIT number by picking
    // digits in order on each line.
    //
    // Key idea:
    //   - Track the biggest digit seen so far (this can be tens).
    //   - Every new digit can pair with that to form a candidate.
    //   - The tens position dominates, so only track the max left digit.
    // ------------------------------------------------------------
    private int SolvePart1(string[] lines)
    {
        int total = 0;

        foreach (string line in lines)
        {
            int bestLeft = 0;   // Best tens digit seen so far
            int bestPair = 0;   // Best two-digit number for this line

            foreach (char c in line)
            {
                int d = c - '0';

                // Form a valid two-digit number using the best left digit
                int candidate = bestLeft * 10 + d;

                if (candidate > bestPair)
                    bestPair = candidate;

                // Update bestLeft so future digits can use it
                if (d > bestLeft)
                    bestLeft = d;
            }

            total += bestPair;
        }

        return total;
    }

    // ------------------------------------------------------------
    // PART 2
    // Pick EXACTLY 12 digits, in order, to create the biggest
    // possible 12-digit number.
    //
    // Classic greedy rule:
    //   - Repeatedly remove the *first* digit where digit[i] < digit[i+1].
    //   - If no such digit exists (the sequence is non-increasing),
    //     remove the final digit.
    //
    // This guarantees the lexicographically largest remaining sequence.
    // ------------------------------------------------------------
    private ulong SolvePart2(string[] lines)
    {
        ulong total = 0;

        foreach (string line in lines)
        {
            string work = line;

            // Keep deleting digits until we're down to 12
            while (work.Length > 12)
            {
                bool removed = false;

                // Find the first "bad" digit and remove it
                for (int i = 0; i < work.Length - 1; i++)
                {
                    if (work[i] < work[i + 1])
                    {
                        work = work.Remove(i, 1);
                        removed = true;
                        break; // Important: remove only ONE digit this round
                    }
                }

                // If nothing was removed, the sequence never increased:
                // ex: 987654321111
                // Remove the last digit in this case.
                if (!removed)
                {
                    work = work[..^1];
                }
            }

            // Now we have exactly 12 digits; accumulate the numeric value
            total += ulong.Parse(work);
        }

        return total;
    }
}
