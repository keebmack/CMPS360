package Swing;

import javax.swing.JOptionPane;

public class SurveyForHW {

    public static void main(String[] args) {
        String[] questions = {
            "Who is the best player for the Steelers?\nA.Mason Cole B.TJ Watt C.Levi Wallace",
            "Who deserves to win DPOY?\nA.TJ Watt B.Myles Garrett C.Maxx Crosby",
            "How many players on the Steelers made the Pro Bowl?\nA.0 B.5 C.3"
        };

        char[] answers = {'B', 'A', 'C'};
        char ans = ' ';
        int x, correct = 0;
        String entry, incorrectQuestions = "";
        boolean isGood;

        for (x = 0; x < questions.length; ++x) {
            isGood = false;
            int firstError = 0;
            while (!isGood) {
                isGood = true;
                entry = JOptionPane.showInputDialog(null, questions[x]);
                ans = entry.charAt(0);
                if (ans != 'A' && ans != 'B' && ans != 'C') {
                    isGood = false;
                    if (firstError == 0) {
                        questions[x] = questions[x] + "\nYour answer must be A, B, or C.";
                        firstError = 1;
                    }
                }
            }
            if (ans == answers[x]) {
                ++correct;
                JOptionPane.showMessageDialog(null, "Correct!");
            } else {
                JOptionPane.showMessageDialog(null, "Incorrect. The correct answer is " + answers[x]);
                incorrectQuestions += (x + 1) + ". " + questions[x] + "\nCorrect Answer: " + answers[x] + "\n";
            }
        }
        JOptionPane.showMessageDialog(null, "You received " + correct + " right and\n" +
            (questions.length - correct) + " wrong.\n\nIncorrect Questions:\n" + incorrectQuestions);
    }
}
