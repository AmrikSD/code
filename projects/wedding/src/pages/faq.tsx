import Accordion from "@/app/Accordion"
import React from "react"

const faqData = [
    {
        title: "When?",
        children:
            <>
                <h3>2nd Feb 2025</h3>
            </>
    },
    {
        title: "Where?",
        children:
            <>
                <h3>Kundala Resort Hua Hin, Thailand</h3>
                <p>122/64 Mooban Khao Takiab</p>
                <p>Nong Kae,</p>
                <p>Hua Hin, Prachuap Khiri Khan, 77110.</p>
                <p><a href="https://www.kundalabeachresort.com/" target="kundala">Visit their website here</a></p>
            </>
    },
    {
        title: "Can I stay in the resort before the wedding?",
        children:
            <>
                <p>Yes.</p>
                <p>
                    While we can&apos;t cover costs for stays outside of the wedding day, we’re excited to offer our guests a special discount at the venue! 
                    If you&apos;d like to stay before or after the wedding night, please let us know so we can help you book at the discounted rate.
                    Booking through us will ensure you get the best price.
                </p>
                <table>
                    <tr>
                        <th>Room</th>
                        <th>🇹🇭THB</th>
                        <th>🇬🇧GBP</th>
                        <th>ℹ️ Info</th>
                    </tr>
                    <tr>
                        <td>Deluxe Room</td>
                        <td>3,500</td>
                        <td>what</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#deluxe" target="kundala-deluxe">Link</a></td>
                    </tr>
                    <tr>
                        <td>Onsen Villa</td>
                        <td>6,000</td>
                        <td>what</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#onsen" target="kundala-onsen">Link</a></td>
                    </tr>
                    <tr>
                        <td>Beachfront Onsen Villa</td>
                        <td>6,800</td>
                        <td>what</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#onsen" target="kundala-beachfront">Link</a></td>
                    </tr>
                    <tr>
                        <td>Private Pool Suite</td>
                        <td>7,600</td>
                        <td>what</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#private" target="kundala-private">Link</a></td>
                    </tr>
                        <td>Lower Beachfront with Hot Jacuzzi</td>
                        <td>8,200</td>
                        <td>what</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#lower" target="kundala-lower">Link</a></td>
                    <tr>
                        <td>Upper Beachfront with Hot Jacuzzi</td>
                        <td>8,900</td>
                        <td>what</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#upper" target="kundala-upper">Link</a></td>
                    </tr>
                    <tr>
                        <td>Sky Pool Suite Room</td>
                        <td>11,000</td>
                        <td>what</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#sky" target="kundala-sky">Link</a></td>
                    </tr>
                    <tr>
                        <td>Kundala Residence Room (4 people)</td>
                        <td>12,000</td>
                        <td>what</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#two" target="kundala-two">Link</a></td>
                    </tr>
                </table>
            </>
    },
    {
        title: "How do we get there?",
        children:
            <>
                <p>No Worries!</p>
                <p>We know it&apos;s a lot to ask for people to travel all the way to Thailand</p>
                <p>We will do something in the UK when we&apos;re back</p>
            </>
    },
    {
        title: "I can't make it",
        children:
            <>
                <p>No Worries!</p>
                <p>We know it&apos;s a lot to ask for people to travel all the way to Thailand</p>
                <p>We will do something in the UK when we&apos;re back</p>
            </>
    },
    {
        title: "I can't make it",
        children:
            <>
                <p>No Worries!</p>
                <p>We know it&apos;s a lot to ask for people to travel all the way to Thailand</p>
                <p>We will do something in the UK when we&apos;re back</p>
            </>
    },
    {
        title: "I can't make it",
        children:
            <>
                <p>No Worries!</p>
                <p>We know it&apos;s a lot to ask for people to travel all the way to Thailand</p>
                <p>We will do something in the UK when we&apos;re back</p>
            </>
    },
]

export default function Faq() {
    return <>
        <div>
            {faqData.map((faq) => (
                <Accordion title={faq.title} key={faq.title}> 
                    {faq.children}
                </Accordion>
            ))}
        </div>
    </>
}
