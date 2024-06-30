package uk.co.amrik.steel.user.api.v1;

import com.google.inject.Inject;
import org.jooq.impl.DSL;
import uk.co.amrik.steel.API.Api;
import uk.co.amrik.steel.persistence.DatabaseService;
import uk.co.amrik.steel.user.model.User;
import java.util.List;
import java.util.Optional;

public class UserService implements Api<User> {

    private final DatabaseService databaseService;

    @Inject
    public UserService(
            DatabaseService databaseService
    ){
        this.databaseService = databaseService;
    }

    @Override
    public List<User> getAll() {
        return databaseService.getDsl().select()
                .from(DSL.table(DSL.name("user")))
                .fetch()
                .map(record -> record.into(User.class));
    }

    @Override
    public Optional<User> get(Integer id) {
        return databaseService.getDsl().select()
                .from(DSL.table(DSL.name("user")))
                .where(DSL.field(DSL.name("id")).eq(id))
                .fetchOptional()
                .map(record -> record.into(User.class));
    }

    @Override
    public void delete(Integer id) {

    }

    @Override
    public User put(User entity) {
        return null;
    }
}
