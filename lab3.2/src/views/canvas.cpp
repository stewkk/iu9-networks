#include "canvas.hpp"

#include <gdkmm.h>

namespace lab32 {
    Canvas::Canvas() {
        set_draw_func(sigc::mem_fun(*this, &Canvas::OnDraw));
    }
    Canvas::~Canvas() {}

    void Canvas::OnDraw(const Cairo::RefPtr<Cairo::Context>& cr, int width, int height) {
        auto style_context = get_style_context();

        style_context->render_background(cr, 0, 0, width, height);

        cr->arc(width / 2.0, height / 2.0, std::min(width, height) / 2.0, 0, 2 * M_PI);

        auto color = style_context->get_color();
        Gdk::Cairo::set_source_rgba(cr, color);

        cr->fill();
    }
}
