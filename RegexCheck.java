import java.util.regex.*;

public class RegexCheck {

    static Pattern rx(int n) {
        return Pattern.compile("a?".repeat(n) + "a".repeat(n));
    }

    public static void main(String[] args) {
        int n = Integer.valueOf(args[0]);
        Pattern p = rx(n);
        String s = "a".repeat(n);
        Matcher m = p.matcher(s);
        m.find();
        System.out.printf("n: %d match: [%d %d]\n", n, m.start(), m.end());
    }
}
