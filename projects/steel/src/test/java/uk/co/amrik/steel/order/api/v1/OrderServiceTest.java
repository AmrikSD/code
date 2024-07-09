package uk.co.amrik.steel.order.api.v1;

import com.google.inject.Guice;
import com.google.inject.Inject;
import com.google.inject.Injector;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import uk.co.amrik.steel.order.persistence.DatabaseExtension;
import uk.co.amrik.steel.persistence.DatabaseService;

import java.util.List;

import static org.assertj.core.api.AssertionsForInterfaceTypes.assertThat;
import static org.junit.jupiter.api.Assertions.assertEquals;

@ExtendWith(DatabaseExtension.class)
public class OrderServiceTest {

    @Inject
    private OrderService underTest;


    @BeforeEach
    public void setUp() {
        Injector injector = Guice.createInjector(
                binder -> {
                    binder.bind(DatabaseService.class);
                    binder.bind(OrderService.class);
                }
        );
        injector.injectMembers(this);
    }

    @Test
    public void emptyDatabaseReturnsNothing(){
        List<OrderResponse> expected = List.of();
        List<OrderResponse> actual = underTest.getAll();

        assertEquals(expected, actual);
    }

    @Test
    public void put(){
        List<OrderResponse> beforePut = underTest.getAll();

        underTest.put(new OrderRequest("test"));
        underTest.put(new OrderRequest("testTwo"));

        List<OrderResponse> afterPut = underTest.getAll();

        assertEquals(List.of(), beforePut);

        assertThat(afterPut).hasSize(2)
                .extracting(OrderResponse::name)
                .containsExactlyInAnyOrder("test", "testTwo");

    }
}