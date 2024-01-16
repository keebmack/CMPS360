import java.util.Scanner;

public class LuckyNumber {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int luckyNumber = 10;

        System.out.println("Welcome to the Lucky Number Game!");
        System.out.print("Guess the lucky number (between 1 and 20): ");

        int userGuess = scanner.nextInt();

        if (userGuess == luckyNumber) {
            System.out.println("Great guess! " + luckyNumber + " is the lucky number.");
        } else {
            System.out.println("Sorry, try again. " + userGuess + " is not the lucky number.");
        }

        scanner.close();
    }
}