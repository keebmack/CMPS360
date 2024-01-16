import java.util.Random;

import javax.swing.JOptionPane;

public class RandomNumber {
    public static void main(String[] args) {
        Random random = new Random();

        int randomNumber = random.nextInt();

        JOptionPane.showMessageDialog(null, "Random Number: " + randomNumber);
    }
}