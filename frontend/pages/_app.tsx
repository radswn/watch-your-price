import { AppProps } from 'next/app';
import { Navbar } from 'components/navbar';
import Head from 'next/head'
import React from 'react';
import 'styles/globals.css';

function MyApp({ Component, pageProps }: AppProps) {
  return <>
    <Head>
        <meta charSet='utf-8' />
        <meta httpEquiv='X-UA-Compatible' content='IE=edge' />
        <meta name='viewport' content='width=device-width,initial-scale=1' />
        <meta name='description' content='Description' />
        <meta name='keywords' content='Keywords' />
        <title>Watch Your Price</title>

        <link rel="manifest" href="/manifest.json" />
        <link href='/favicon-16x16.png' rel='icon' type='image/png' sizes='16x16' />
        <link href='/favicon-32x32.png' rel='icon' type='image/png' sizes='32x32' />
        <link href="/apple-icon.png" rel="apple-touch-icon"></link>
        <meta name="theme-color" content="#317EFB" />
    </Head>
    <Navbar/>
    <div>
      <Component {...pageProps} />
    </div>
  </>
}

export default MyApp
