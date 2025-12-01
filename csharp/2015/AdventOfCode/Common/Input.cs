namespace Common;

public static class Input
{
    public static string[] Lines(string path)
        => File.ReadAllLines(path);

    public static string Text(string path)
        => File.ReadAllText(path);
}
