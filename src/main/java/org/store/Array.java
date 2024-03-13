package org.store;

public class Array {
    private String[][] store;
    private int[] cursor;

    public Array() {
        this.cursor = new int[]{0, 0};
        this.store = new String[500][500];
    }

    public void add(String text) {
        // todo
        store[cursor[0]][cursor[1]] = text;
    }

    public void delete(int row, int col) {
        // todo
    }
}
