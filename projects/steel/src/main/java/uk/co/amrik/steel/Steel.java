package uk.co.amrik.steel;

import com.google.inject.Guice;

public class Steel {

    public static void main(String ...args) throws Exception {
        Guice.createInjector(
                new SteelModule()
        ).getInstance(Server.class).run();
    }


}
