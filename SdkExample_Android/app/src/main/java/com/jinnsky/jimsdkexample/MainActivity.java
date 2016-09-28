package com.jinnsky.jimsdkexample;

import android.speech.tts.TextToSpeech;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.widget.TextView;

import go.jimsdk.*;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        TextView originTextView = (TextView) findViewById(R.id.originContentView);
        TextView urlTextView = (TextView) findViewById(R.id.urlContentView);

        try {
            Client client = Jimsdk.newClient();
            originTextView.setText(client.getOrigin());
            urlTextView.setText(client.getURL());
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
