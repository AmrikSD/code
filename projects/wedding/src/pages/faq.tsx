import Accordion from "@/app/Accordion"
import React from "react"

const faqData = [
    {
        title: "What date should I RSVP by?",
        children:
            <>
                <h2>11th August 2024</h2>
                <p>We would really appreciate if you could officially RSVP using the form a the bottom of this page before this date.</p>
            </>
    },
    {
        title: "When is the wedding?",
        children:
            <>
                <h2>2nd February 2025</h2>
            </>
    },
    {
        title: "Where is the wedding?",
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
                    While we can&apos;t cover costs for stays outside of the wedding night, the venue will offer discounted rates for wedding guests! 
                    If you&apos;d like to stay before or after the wedding night, please let us know so we can help you book at the discounted rate.
                </p>
                <table>
                    <tr>
                        <th>Room</th>
                        <th>🇹🇭THB</th>
                        <th>ℹ️ Info</th>
                    </tr>
                    <tr>
                        <td>Deluxe Room</td>
                        <td>3,500</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#deluxe" target="kundala-deluxe">Link</a></td>
                    </tr>
                    <tr>
                        <td>Onsen Villa</td>
                        <td>6,000</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#onsen" target="kundala-onsen">Link</a></td>
                    </tr>
                    <tr>
                        <td>Beachfront Onsen Villa</td>
                        <td>6,800</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#onsen" target="kundala-beachfront">Link</a></td>
                    </tr>
                    <tr>
                        <td>Private Pool Suite</td>
                        <td>7,600</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#private" target="kundala-private">Link</a></td>
                    </tr>
                        <td>Lower Beachfront with Hot Jacuzzi</td>
                        <td>8,200</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#lower" target="kundala-lower">Link</a></td>
                    <tr>
                        <td>Upper Beachfront with Hot Jacuzzi</td>
                        <td>8,900</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#upper" target="kundala-upper">Link</a></td>
                    </tr>
                    <tr>
                        <td>Sky Pool Suite Room</td>
                        <td>11,000</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#sky" target="kundala-sky">Link</a></td>
                    </tr>
                    <tr>
                        <td>Kundala Residence Room (4 people)</td>
                        <td>12,000</td>
                        <td><a href="https://www.kundalabeachresort.com/index.php?page=accommodation_2#two" target="kundala-two">Link</a></td>
                    </tr>
                </table>
            </>
    },
    {
        title: "How do we get there?",
        children:
            <>
                <p>Our venue is in Hua Hin, a district a couple of hours&apos; drive from Bangkok.</p>
                <p>The nearest international airport is Suvarnabhumi Airport (BKK).</p>
                <p>We will arrange transportation for UK guests from Bangkok straight to the venue on the day.</p>
                <p>If you want to go stay around the area before the wedding day, we can also recommend ways to get there.</p>
            </>
    },
    {
        title: "What is the dress code?",
        children:
            <>
                <p>Our wedding will be taking place on the beach.</p>
                <p>Please wear dressy casual to semi-formal attire so we can all look great in photos while staying comfortable!</p>
                <p>Women: heels will be impractical in the sand, consider flats or wedges</p>
                <p>Men: blazers and ties are optional and may be uncomfortable in the Thai heat.</p>
            </>
    },
    {
        title: "I can't make it",
        children:
            <>
                <p>No Worries!</p>
                <p>Please still fill out the RSVP Below and select &quot;no&quot;, though!</p>
                <p>We know it&apos;s a lot to ask for people to travel all the way to Thailand.</p>
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
