package uk.co.amrik.steel.order.api.v1;

import com.google.inject.Inject;
import org.jooq.impl.DSL;
import uk.co.amrik.steel.API.Api;
import uk.co.amrik.steel.order.model.Order;
import uk.co.amrik.steel.persistence.DatabaseService;

import java.util.List;
import java.util.Optional;

public class OrderService implements Api<Order> {

    private final DatabaseService databaseService;

    @Inject
    public OrderService(DatabaseService databaseService) {
        this.databaseService = databaseService;
    }

    @Override
    public List<Order> getAll() {
        return databaseService.getDsl().select()
                .from(DSL.table(DSL.name("order")))
                .fetch()
                .map(record -> record.into(Order.class));
    }

    @Override
    public Optional<Order> get(Integer id) {
        return Optional.empty();
    }

    @Override
    public void delete(Integer id) {

    }

    @Override
    public Order put(Order entity) {
        return null;
    }
}
