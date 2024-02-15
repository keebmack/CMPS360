using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;

namespace Module_5.Pages;

public class SteelersModel : PageModel
{
    private readonly ILogger<SteelersModel> _logger;

    public SteelersModel(ILogger<SteelersModel> logger)
    {
        _logger = logger;
    }

    public void OnGet()
    {

    }
}
