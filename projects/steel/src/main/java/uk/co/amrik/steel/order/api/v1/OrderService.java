package uk.co.amrik.steel.order.api.v1;

import com.google.inject.Inject;
import org.jooq.impl.DSL;
import uk.co.amrik.steel.API.Api;
import uk.co.amrik.steel.persistence.DatabaseService;

import java.util.List;
import java.util.Optional;

public class OrderService implements Api<OrderRequest, OrderResponse> {

    private final DatabaseService databaseService;

    @Inject
    public OrderService(DatabaseService databaseService) {
        this.databaseService = databaseService;
    }

    @Override
    public List<OrderResponse> getAll() {
        return databaseService.getDsl().select()
                .from(DSL.table(DSL.name("order")))
                .fetch()
                .map(record -> record.into(OrderResponse.class));
    }

    @Override
    public Optional<OrderResponse> get(Integer id) {
        return databaseService.getDsl().select()
                .from(DSL.table(DSL.name("order")))
                .where(DSL.field(DSL.name("id")).eq(id))
                .fetchOptional()
                .map(record -> record.into(OrderResponse.class));
    }

    @Override
    public int delete(Integer id) {
        return databaseService.getDsl().delete(DSL.table(DSL.name("order")))
                .where(DSL.field(DSL.name("id")).eq(id))
                .execute();
    }

    @Override
    public Optional<OrderResponse> put(OrderRequest entity) {
        return databaseService.getDsl().insertInto(DSL.table(DSL.name("order")))
                .set(DSL.field(DSL.name("name"), String.class), entity.name())
                .returning()
                .fetchOptionalInto(OrderResponse.class);
    }
}
