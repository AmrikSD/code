package uk.co.amrik.steel.permissions;

import com.google.inject.Inject;
import org.jooq.impl.DSL;
import uk.co.amrik.steel.persistence.DatabaseService;

import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class PermissionsService {

    private final DatabaseService databaseService;

    @Inject
    PermissionsService(
            DatabaseService databaseService
    ){
        this.databaseService = databaseService;
    }

    public List<Role> getRoles(){
        return databaseService.getDsl().select(DSL.field("name"))
                .from(DSL.table("role"))
                .stream()
                .map(r -> r.get("name", String.class))
                .map(String::toUpperCase)
                .map(Role::valueOf)
                .toList();
    }

    public void bootstrapRoles() {
        Set<String> enumRoles = Stream.of(Role.values())
                .map(Enum::name)
                .collect(Collectors.toSet());


        Set<String> dbRoles = databaseService.getDsl().select(DSL.field("name"))
                .from(DSL.table("role"))
                .stream()
                .map(r -> r.get("name", String.class).toUpperCase())
                .collect(Collectors.toSet());

        Set<String> rolesNotInDb = new HashSet<>(enumRoles);
        rolesNotInDb.removeAll(dbRoles);

        if (!rolesNotInDb.isEmpty()){
            rolesNotInDb.forEach(role ->
                    databaseService.getDsl().insertInto(DSL.table("role"))
                            .columns(DSL.field("name"))
                            .values(role)
                            .onConflict(DSL.field("name")).doNothing()
                            .execute()
            );
        }

    }

}
