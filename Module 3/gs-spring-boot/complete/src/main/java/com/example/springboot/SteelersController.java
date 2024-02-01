package com.example.springboot;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class SteelersController {

    @GetMapping("/steelers/founded")
    public String steelersFounded() {
        int foundedYear = 1933;
        if (foundedYear == 1933) {
            return "The Pittsburgh Steelers were founded in 1933.";
        } else {
            return "Founded in a different year.";
        }
    }

    @GetMapping("/steelers/superbowls")
    public String steelersSuperBowls() {
        int superBowlWins = 6;
        if (superBowlWins == 6) {
            return "The Steelers have won 6 Super Bowls.";
        } else {
            return "A different number of Super Bowl wins.";
        }
    }

    @GetMapping("/steelers/stadium")
    public String steelersStadium() {
        String stadiumName = "Acrisure Stadium";
        if ("Acrisure Stadium".equals(stadiumName)) {
            return "The Steelers play at Acrisure Stadium.";
        } else {
            return "They play at a different stadium.";
        }
    }

    @GetMapping("/steelers/mascot")
    public String steelersMascot() {
        String mascotName = "Steely McBeam";
        if ("Steely McBeam".equals(mascotName)) {
            return "The Steelers' mascot is Steely McBeam.";
        } else {
            return "A different mascot name.";
        }
    }
}
