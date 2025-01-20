import React from "react"

import Section from "@app/section"
import Start from "@pages/start"
import Location from "@pages/location"
import Faq from "@pages/faq"
import Rsvp from "@pages/rsvp"

import img from "@public/image (1).jpg"

export default function Index() {
    return <>
        <Section name="start"> 
            <Start />
        </Section>
        <Section name="location"> 
            <Location />
        </Section>
        <Section name="faq"> 
            <Faq />
        </Section>
        <Section name="rsvp"> 
            <Rsvp />
        </Section>
        <Section name="kate"> 
            <img src={img.src} alt="kate and amrik"/>
            <h2>Made with ❤️ and 🫖 and 🍫 by 🥚 and 🍅</h2>
        </Section>
    </>
}
