package com.sim.hummus;

import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by ariel on 16.08.2017.
 */
class FileInterpreter {
    void interpreteAllLine(String arg, Main m) {
        List<String> lines = new ArrayList<>();
        try {
            lines = Files.readAllLines(Paths.get(new File(arg).toURI()));
        } catch (Exception e) {
            System.out.println(e.getMessage());
            try {
                System.in.read();
            } catch (Exception ignored) {
            }
            System.exit(-1);
        }

        for (String line : lines) {
            Pair<String,Boolean> result = m.interpretLine(line);

            if (result.getValue()){
                System.out.println(result.getKey());
            }
        }
    }
}