using System;

namespace MyConsoleApp
{
    class Program
    {
        static void Main(string[] args)
        {
            Random random = new Random();
            int firstDiceRoll = random.Next(1, 21);
            int firstSkillCheck = 10;
            int secondSkillCheck = 15;

            Console.WriteLine("Rolling 1d20 for the first skill check...");
            Console.WriteLine($"Result: {firstDiceRoll}");

            if (firstDiceRoll == 20)
            {
                Console.WriteLine("Critical Success on the first check! Automatically successful on the second check.");
            }
            else if (firstDiceRoll == 1)
            {
                Console.WriteLine("Critical Failure!");
            }
            else if (firstDiceRoll >= firstSkillCheck)
            {
                Console.WriteLine("Success on the first check! Proceeding to the second skill check...");

                int secondDiceRoll = random.Next(1, 21);
                Console.WriteLine($"Rolling 1d20 for the second skill check...");
                Console.WriteLine($"Result: {secondDiceRoll}");

                if (secondDiceRoll == 20)
                {
                    Console.WriteLine("Critical Success on the second check!");
                }
                else if (secondDiceRoll == 1)
                {
                    Console.WriteLine("Critical Failure on the second check!");
                }
                else if (secondDiceRoll >= secondSkillCheck)
                {
                    Console.WriteLine("Success on the second check!");
                }
                else
                {
                    Console.WriteLine("Failure on the second check!");
                }
            }
            else
            {
                Console.WriteLine("Failure on the first check!");
            }
        }
    }
}