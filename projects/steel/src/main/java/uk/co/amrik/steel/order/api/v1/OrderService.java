package uk.co.amrik.steel.order.api.v1;

import uk.co.amrik.steel.order.model.Order;
import java.util.List;

public class OrderService implements OrderResource {
    @Override
    public List<Order> order() {
        return List.of(
                new Order("hi")
        );
    }
}
