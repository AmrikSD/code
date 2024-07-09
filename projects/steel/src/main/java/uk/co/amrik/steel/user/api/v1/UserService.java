package uk.co.amrik.steel.user.api.v1;

import com.google.inject.Inject;
import org.jooq.impl.DSL;
import uk.co.amrik.steel.API.Api;
import uk.co.amrik.steel.persistence.DatabaseService;

import java.util.List;
import java.util.Optional;

public class UserService implements Api<UserRequest, UserResponse> {

    private final DatabaseService databaseService;

    @Inject
    public UserService(
            DatabaseService databaseService
    ){
        this.databaseService = databaseService;
    }

    @Override
    public List<UserResponse> getAll() {
        return databaseService.getDsl().select()
                .from(DSL.table(DSL.name("user")))
                .fetch()
                .map(record -> record.into(UserResponse.class));
    }

    @Override
    public Optional<UserResponse> get(Integer id) {
        return databaseService.getDsl().select()
                .from(DSL.table(DSL.name("user")))
                .where(DSL.field(DSL.name("id")).eq(id))
                .fetchOptional()
                .map(record -> record.into(UserResponse.class));
    }

    @Override
    public int delete(Integer id) {
        return 0;
    }

    @Override
    public Optional<UserResponse> put(UserRequest entity) {
        return Optional.empty();
    }
}
