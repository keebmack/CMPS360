using Microsoft.AspNetCore.Mvc.RazorPages;

namespace Module_5.Pages;

public class AboutModel : PageModel
{
    public string Message { get; set; } = string.Empty; // Initialize to ensure non-null
    
    public void OnGet()
    {
        if (DateTime.Now.DayOfWeek == DayOfWeek.Sunday)
        {
            Message = "It's a Steelers Sunday!";
        }
        else
        {
            Message = "Boo, no football today";
        }
    }
}
