package Swing;

import javax.swing.JOptionPane;
public class Survey {

    public static void main(String[] args){
        String[]questions =
        {
            "How many players played well this season?\nA.4 B.7 C.9",
            "How many points did Boswell score this season?\nA.3 B.40 C.114",
            "How many penalties did the Steelers have in game 3 of the season?\nA.1 B.3 C.0"
        };

        char[] answers = {'C','C','B'};
        char ans = ' ';
        int x, correct = 0;
        String entry;
        boolean isGood;

        for(x = 0; x < questions.length; ++x)
        {
            isGood = false;
            int firstError = 0;
            while(!isGood)
            {
                isGood = true;
                entry = JOptionPane.showInputDialog(null,questions[x]);
                ans = entry.charAt(0);
                if(ans != 'A' && ans != 'B' && ans != 'C')
                {
                    isGood = false;
                    if(firstError == 0)
                    {
                        questions[x] = questions[x] +
                        "\nYour answer must be A, B or C.";
                        firstError = 1;
                    }
                }
            }
            if(ans == answers[x])
            {
                ++correct;
                JOptionPane.showMessageDialog(null, "Correct!");
            }
            else
                JOptionPane.showMessageDialog(null, "The correct answer is "
                + answers[x]);
        }
        JOptionPane.showMessageDialog(null, "You received "
        + correct + " right and\n" +(questions.length -correct) + " wrong");
    }
}
