import type { ActionFunction, ActionFunctionArgs, LoaderFunction } from "@remix-run/cloudflare";
import { json } from "@remix-run/cloudflare";
import CountDown from "./Countdown";
import { Form, useLoaderData } from "@remix-run/react";
import { PropsWithChildren, ReactNode, ReactPropTypes, RefObject, useEffect, useRef } from "react";

interface Env {
  MY_DB: D1Database;
}

type rsvp = {
    rsvp_id:number,
    first_name:number,
    last_name:number,
    plus_one:number,
    plus_one_first_name:number,
    plus_one_last_name:number,
    coming_to_thailand:number,
    coming_to_uk:number,
    diet:number
}

export const loader: LoaderFunction = async ({ context })  => {
  const env = context.cloudflare.env as Env
  const { results } = await env.MY_DB.prepare("SELECT * FROM rsvp ORDER BY rowid desc LIMIT 99").all(); //TODO: remove LIMIT
  return json(results);
};

export const action: ActionFunction = async ({request, context} : ActionFunctionArgs) => {
    const env = context.cloudflare.env

    const formData = await request.formData()
    const first_name = "first_name"
    const last_name = "last_name"
    const plus_one = 1
    const plus_one_first_name = "plus_one_first_name"
    const plus_one_last_name = "plus_one_last_name"
    const coming_to_thailand = 2
    const coming_to_uk = 2
    const diet = "diet"

    console.log(formData)

    return await env.MY_DB.prepare(`
        INSERT INTO rsvp (first_name, last_name, plus_one, plus_one_first_name, plus_one_last_name, coming_to_thailand, coming_to_uk, diet)
        VALUES (?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8);
    `).bind(
        first_name,
        last_name,
        plus_one,
        plus_one_first_name,
        plus_one_last_name,
        coming_to_thailand,
        coming_to_uk,
        diet,
    ).run().then((resp) => {
        return resp
    })
}

const Page = ({name, children}: PropsWithChildren<{name?: string}>): ReactNode => {
    const ref = useRef<HTMLDivElement>(null);

    return (
        <section key={name} data-page={name}>
            <div className="container" ref={ref}>
                {children}
            </div>
        </section>
    )
}
const scrollPageIntoView = (name: string) => {
    const element = document.querySelector(`section[data-page='${name}']`);
    if (element) {
        element.scrollIntoView({ behavior: 'smooth' });
    }
};

const NavButton = ({to, text}: {to: string, text: string}): ReactNode => {
    return (
        <button type="button" onClick={() => scrollPageIntoView(to)}>{text}</button>
    )
}


export default function Index() {

  const results: rsvp[] = useLoaderData<typeof loader>() as rsvp[];

  return (
    <>
        <Page name="landing">
            <h1>We&apos;re getting married!</h1>
            <h1>You&apos;re invited!</h1>
            <div>
                <p>We (<u className="kate">Kate</u> &amp; Amrik) would love to invite you to our wedding!</p>
                <p>For more info, just scroll down!</p>
                <p>To <b>RSVP</b> just scroll down, or <b>TEXT</b> us!</p>
            </div>
            <div>
                <NavButton to="countdown" text="More Info"/>
                <NavButton to="rsvp-1" text="RSVP"/>
            </div>
        </Page>
        <Page name="countdown">
            <h1><a href="https://www.google.com/maps/place/Hua+Hin,+Hua+Hin+District,+Prachuap+Khiri+Khan+77110,+Thailand/@12.5913105,99.9409497,13z/data=!3m1!4b1!4m6!3m5!1s0x30fdaaf400997211:0x40223bc2c381260!8m2!3d12.5683747!4d99.9576888!16s%2Fg%2F123dmcgrf?entry=ttu">Hua Hin</a>, Thailand</h1>
            <CountDown target={new Date("2025-02-02")}/>
            <p>yes, a destination wedding :)</p>
            <p>Don&apos;t worry, there will be another celebration in the UK</p>
            <NavButton to="rsvp-1" text="Next Page"/>
        </Page>
        <Form method="POST">
            <Page name="rsvp-1">
                <fieldset>
                    <legend>RSVP</legend>
                    <div>
                        <div className="form-group names">
                            <input type="text" placeholder="Name" required />
                            <input type="text" placeholder="Surname" required />
                        </div>
                        <div className="form-group coming-thailand">
                          <div className="switch">	
                            <input type="radio" name="choice" id="yes" defaultChecked />
                            <label htmlFor="yes">See you there</label>
                            <input type="radio" name="choice" id="no" />
                            <label htmlFor="no">Sorry, can&apos; make it!</label>
                            <span className="switchFilter"></span>
                          </div>
                        </div>
                        <button type="button" onClick={() => scrollPageIntoView("landing")}>Back to start??</button>
                    </div>
                </fieldset>
            </Page>
            <Page>
                <p>part 2</p>
                <button type="submit">RSVP!</button>
            </Page>
        </Form>
        {
            results.map((result: rsvp) => {
                return (
                    <Page key={result.rsvp_id}>
                        <h1>{result.first_name}</h1>
                        <ul>
                            <li>Name: {result.first_name}</li>
                            <li>Surname: {result.last_name}</li>
                            <li>Plus One: {result.plus_one}</li>
                            <li>Coming Thailand: {result.coming_to_thailand}</li>
                            <li>Coming UK {result.coming_to_uk}</li>
                            <li>PO Name: {result.plus_one_first_name}</li>
                            <li>PO Surname: {result.plus_one_last_name}</li>
                        </ul>
                        <p>{result.diet}</p>
                    </Page>
                )
            })
        }
        <Page name="footer">
            <p>Website created with ❤️ and 🫖 by <u className="amrik">Amrik</u></p>
        </Page>
    </>
  );
}
