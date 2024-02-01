package com.example.springboot;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;

@SpringBootTest
@AutoConfigureMockMvc
public class SteelersControllerTest {

    @Autowired
    private MockMvc mvc;

    @Test
    public void testSteelersFounded() throws Exception {
        mvc.perform(get("/steelers/founded").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string("The Pittsburgh Steelers were founded in 1933."));
    }

    @Test
    public void testSteelersSuperBowls() throws Exception {
        mvc.perform(get("/steelers/superbowls").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string("The Steelers have won 6 Super Bowls."));
    }

    @Test
    public void testSteelersStadium() throws Exception {
        mvc.perform(get("/steelers/stadium").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string("The Steelers play at Acrisure Stadium."));
    }

    @Test
    public void testSteelersMascot() throws Exception {
        mvc.perform(get("/steelers/mascot").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string("The Steelers' mascot is Steely McBeam."));
    }
}
