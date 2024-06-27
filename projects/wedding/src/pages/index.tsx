import Section from "@/app/section"
import React from "react"
import Start from "./start"
import Location from "./location"
import Faq from "./faq"
import Rsvp from "./rsvp"

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
    </>
}
