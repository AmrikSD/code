package uk.co.amrik.steel.user.api.v1;

import com.google.inject.Inject;
import jakarta.ws.rs.WebApplicationException;
import jakarta.ws.rs.core.Response;
import org.jooq.impl.DSL;
import uk.co.amrik.steel.persistence.DatabaseService;
import uk.co.amrik.steel.user.model.User;
import java.util.List;
import java.util.Optional;

public class UserService implements UserResource {

    private final DatabaseService databaseService;

    @Inject
    public UserService(
            DatabaseService databaseService
    ){
        this.databaseService = databaseService;
    }

    public List<User> users(){
        return databaseService.getDsl().select()
                .from(DSL.table(DSL.name("user")))
                .fetch(r -> {
                    Integer userId = r.get(DSL.field(DSL.name("id")), Integer.class);
                    String name = r.get(DSL.field(DSL.name("name")), String.class);
                    return new User(userId, name);
                });
    }

    public Optional<User> user(Integer id){
        Optional<User> userOptional = databaseService.getDsl().select()
                .from(DSL.table(DSL.name("user")))
                .where(DSL.field(DSL.name("id")).eq(id))
                .fetchOptional(r -> {
                    Integer userId = r.get(DSL.field(DSL.name("id")), Integer.class);
                    String name = r.get(DSL.field(DSL.name("name")), String.class);
                    return new User(userId, name);
                });

        if (userOptional.isEmpty()){
            throw new WebApplicationException(Response.Status.NOT_FOUND);
        }

        return userOptional;
    }
}
