package com.jinnsky.jimsdkexample;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.widget.TextView;

import go.jimsdk.*;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        TextView contentTextView = (TextView) findViewById(R.id.contentTView);

        try {
            Client client = Jimsdk.newClient("http://api2.jimyun.com", "", "", "");
            contentTextView.setText(Long.toString(client.getServerTimestamp()));
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
