package com.example.computation.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;

@RestController
public class ComputationController {

    @GetMapping("/api/compute")
    public Map<String, Object> compute(@RequestParam double value) {
        // Perform advanced computation
        double result = Math.pow(value, 2) + Math.sqrt(value) * 1.5;

        // Return the result
        Map<String, Object> response = new HashMap<>();
        response.put("inputValue", value);
        response.put("computedValue", result);
        response.put("message", "Computation completed successfully");
        return response;
    }
}
