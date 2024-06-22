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
        title: "I can't make it",
        children:
            <>
                <p>No Worries!</p>
                <p>We know it&apos;s a lot to ask for people to travel all the way to Thailand</p><p>We will do something in the UK when we&apos;re back</p>
            </>
    },
]

export default function Faq() {
    return <>
        <div className="TODO">
            {faqData.map((faq) => (
                <Accordion title={faq.title} key={faq.title}> 
                    {faq.children}
                </Accordion>
            ))}
        </div>
    </>
}
