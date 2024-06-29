package uk.co.amrik.steel.order.api.v1;

import com.google.inject.Inject;
import org.jooq.impl.DSL;
import uk.co.amrik.steel.order.model.Order;
import uk.co.amrik.steel.persistence.DatabaseService;

import java.util.List;

public class OrderService implements OrderResource {

    private final DatabaseService databaseService;

    @Inject
    public OrderService(DatabaseService databaseService) {
        this.databaseService = databaseService;
    }

    @Override
    public List<Order> order() {

        List<Order> orders = databaseService.getDsl().select()
                .from(DSL.table("public.user"))
                .where(
                        DSL.field("id").le(200)
                ).fetch()
                .map(record -> record.into(Order.class));

        return orders;
    }
}
